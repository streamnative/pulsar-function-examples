#[allow(unused_imports)]
use wasmedge_bindgen::*;
use wasmedge_bindgen_macro::*;

use serde::{Deserialize, Serialize};

#[derive(Debug, Deserialize,  Serialize)]
pub struct Student {
    pub name: Option<String>,
    pub age: Option<i32>,
    pub grade: Option<i32>
}

#[wasmedge_bindgen]
pub fn process(mut s: Vec<u8>) -> Vec<u8> {
    s.push(b'!');
    return s
}

#[wasmedge_bindgen]
pub fn process_json(s: Vec<u8>) -> Vec<u8> {
    let stu = serde_json::from_slice::<Student>(&s[..]).unwrap();
    let stu = Student {
        grade: stu.grade.map(|grade| grade + 1),
        ..stu
    };
    let res = serde_json::to_vec(&stu).unwrap();
    return res
}
