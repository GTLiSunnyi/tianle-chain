extern crate serde;
extern crate serde_yaml;
use serde::{Deserialize, Serialize};

// env.yaml 中定义的是链运行时需要的参数

// 定义 config 类型
#[derive(Serialize, Deserialize)]
pub struct Config {
    pub chain_id: String,
    pub algo: String,
    pub home_path: String,
}

impl Config {
    // 读取配置文件
    pub fn new() -> Config {
        let yaml_str = include_str!("../env.yaml");
        serde_yaml::from_str(yaml_str).expect("yaml read failed!")
    }
}
