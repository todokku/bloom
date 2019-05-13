#[macro_use]
extern crate diesel;

#[macro_use]
extern crate diesel_enum_derive;

pub mod validators;
pub mod api;
pub mod domain;
pub mod controllers;

pub const REPORT_MAX_SIZE: u64 = 100_000_000; // 100 MB
