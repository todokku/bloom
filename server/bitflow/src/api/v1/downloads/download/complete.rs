use kernel::{
    api,
    log::macros::*,
    api::middlewares::{
        GetRequestLogger,
        GetRequestId,
        GetRequestAuth,
    },
    KernelError,
};
use futures::{
    future::Future,
    future::ok,
    future::Either,
};
use actix_web::{
    web, Error, HttpRequest, HttpResponse, ResponseError,
};
use crate::{
    controllers,
    domain::download,
};


pub fn post(download_id: web::Path<(uuid::Uuid)>, download_data: web::Json<download::CompleteData>, state: web::Data<api::State>, req: HttpRequest)
-> impl Future<Item = HttpResponse, Error = Error> {    let logger = req.logger();
    let auth = req.request_auth();
    let request_id = req.request_id().0;

    if auth.session.is_none() || auth.account.is_none() {
        return Either::A(ok(KernelError::Unauthorized("Authentication required".to_string()).error_response()));
    }

    return Either::B(
        state.db
        .send(controllers::CompleteDownload{
            download_id: download_id.into_inner(),
            complete_data: download_data.clone(),
            s3_bucket: state.config.s3_bucket(),
            s3_client: state.s3_client.clone(),
            // actor_id: auth.account.expect("error unwraping non none account").id,
            // session_id: auth.session.expect("error unwraping non none session").id,
            request_id,
        })
        .map_err(|_| KernelError::ActixMailbox)
        .from_err()
        .and_then(move |res| {
            match res {
                Ok(_) => {
                    let res = api::Response::data(api::NoData{});
                    ok(HttpResponse::Ok().json(&res))
                },
                Err(err) => {
                    slog_error!(logger, "{}", err);
                    ok(err.error_response())
                },
            }
        })
    );
}
