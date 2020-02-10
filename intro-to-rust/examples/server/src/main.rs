#[macro_use]
extern crate warp;
extern crate chrono;

use chrono::offset::Local;
use warp::Filter;

const PORT: u16 = 3030;

fn main() {
    let index = warp::get2().and(warp::fs::dir("./"));

    let time = path!("time").map(|| {
        let current_time = Local::now();
        format!("The current time in your timezone is: {}", current_time)
    });

    // Combine routes
    let routes = index
        .or(time);

    println!("Listening on port: {}", PORT);
    warp::serve(routes).run(([127, 0, 0, 1], PORT));
}
