// #[global_allocator]
// static ALLOC: snmalloc_rs::SnMalloc = snmalloc_rs::SnMalloc;

use actix_http::{HttpService, KeepAlive};
use actix_service::map_config;
use actix_web::dev::{AppConfig, Body, Server};
use actix_web::http::StatusCode;
use actix_web::{web, App, HttpResponse};
use bytes::Bytes;
use std::env;
use std::thread;
use std::time::Duration;

static mut sleep_time: u32 = 0;

async fn hello() -> HttpResponse {
    unsafe {
        if sleep_time > 0 {
            thread::sleep(Duration::from_millis(sleep_time as u64));
        }
    }

    let res = HttpResponse::with_body(
        StatusCode::OK,
        Body::Bytes(Bytes::from_static(b"Hello, World!")),
    );
    res
}

#[actix_rt::main]
async fn main() -> std::io::Result<()> {
    let mut port: u32 = 8080;

    let args: Vec<String> = env::args().collect();
    if args.len() > 1 {
        unsafe {
            sleep_time = args[1].parse::<u32>().unwrap();
        }
    }
    if args.len() > 2 {
        port = args[1].parse::<u32>().unwrap();
    }

    let addr = format!("0.0.0.0:{}", port);

    // start http server
    Server::build()
        .backlog(1024)
        .bind("benchmark", addr, || {
            HttpService::build()
                .keep_alive(KeepAlive::Os)
                .h1(map_config(
                    App::new().service(web::resource("/hello").to(hello)),
                    |_| AppConfig::default(),
                ))
                .tcp()
        })?
        .start()
        .await
}
