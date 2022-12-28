// use scraper::{Html, Selector};

use reqwest::{header::USER_AGENT, Client};
/*

let client = reqwest::Client::new();
let res = client
    .get("https://www.rust-lang.org")
    .header(USER_AGENT, "My Rust Program 1.0")
    .send()
    .await?; */

 fn main() {
    // let url = "https://npidb.org/doctors/dental_health/dental-assistant_126800000x/1235287046.aspx";
    let url = "https://npidb.org/";
    let client = reqwest::Client::new();
    let res = getResponse(&url, &client);

    // let response = client.get(url).header(USER_AGENT, "Thing!").send();
    // let response = reqwest::blocking::get(url).expect("Could not load url.");
    // let body = response.text().unwrap();

    print!("", res);

    // let document = Html::parse_document(&body);
}

async fn getResponse(url: &str, client: &Client) -> String {
    let res = client.get(url).header(USER_AGENT, "thing").send().await;
    let text = res.unwrap().text().await.unwrap();
    return text
    // Ok(res);
    // return res.text().await?;

    // res.await.unwrap();
}

// Print a web page onto stdout
// fn main() {
//     let url = "https://books.toscrape.com/";
//     let response = reqwest::blocking::get(url).expect("Could not load url.");
//     let body = response.text().unwrap();
//     print!("{}", body);

//     let document = Html::parse_document(&body);

//     let book_selector = Selector::parse("article.product_pod").unwrap();
//     let book_name_selector = Selector::parse("h3 a").unwrap();
//     let book_price_selector = Selector::parse(".price_color").unwrap();
//     // Now the selector is ready to be used. Add the following lines to the main function:
//     for element in document.select(&book_selector) {
//         let book_name_element = element
//             .select(&book_name_selector)
//             .next()
//             .expect("Could not select book name.");
//         let book_name = book_name_element
//             .value()
//             .attr("title")
//             .expect("Could not find title attribute.");
//         let price_element = element
//             .select(&book_price_selector)
//             .next()
//             .expect("Could not find price");
//         let price = price_element.text().collect::<String>();

//         println!("{:?} - {:?}", book_name, price);
//     }
// }
