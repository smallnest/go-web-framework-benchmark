use futures::future;
use std::env;
use std::thread;
use std::time::Duration;
use tokio_minihttp::{Http, Request, Response};
use tokio_proto::TcpServer;
use tokio_service::Service;

struct HelloWorld {
    sleep_time: u32,
}

impl Service for HelloWorld {
    type Request = Request;
    type Response = Response;
    type Error = std::io::Error;
    type Future = future::Ok<Response, std::io::Error>;

    fn call(&self, req: Request) -> Self::Future {
        let mut resp = Response::new();
        match req.path() {
            "/hello" => {
                resp.body("Hello, world!");
                if self.sleep_time > 0 {
                    thread::sleep(Duration::from_millis(self.sleep_time as u64));
                }
            }
            _ => {
                resp.status_code(404, "Not Found");
            }
        }

        future::ok(resp)
    }
}

fn main() {
    let mut sleep_time: u32 = 0;
    let mut port: u32 = 8080;

    let args: Vec<String> = env::args().collect();
    if args.len() > 1 {
        sleep_time = args[1].parse::<u32>().unwrap();
    }
    if args.len() > 2 {
        port = args[1].parse::<u32>().unwrap();
    }

    let addr = format!("0.0.0.0:{}", port).parse().unwrap();
    let mut srv = TcpServer::new(Http, addr);
    srv.threads(num_cpus::get());
    srv.serve(move || Ok(HelloWorld { sleep_time }))
}
