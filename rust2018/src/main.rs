use futures::executor::block_on;

fn main() -> Result<(), Box<dyn std::error::Error>> {
    println!("--- Rust 2018 Demo ---");

    // 1. async / .await
    // 非同期処理のファーストクラスサポート
    let future = say_hello();
    block_on(future);

    // 2. Non-Lexical Lifetimes (NLL)
    // 借用期間を「スコープ」ではなく「実際の使用」で判断
    let mut x = 5;
    let y = &x;
    println!("NLL: y is {}", y); // ここで y のライフタイムは終了
    x = 6; // 2015 までは y のスコープが main 末尾まで続き、ここでエラーだった
    println!("NLL: x updated to {}", x);

    // 3. dyn Trait
    // トレイトオブジェクトであることを dyn キーワードで明示
    let animal: Box<dyn Animal> = Box::new(Dog);
    animal.speak();
    println!("Debug: {:?}", Dog);

    Ok(())
}

async fn say_hello() {
    println!("Async: hello from async function!");
}

trait Animal {
    fn speak(&self);
}

#[derive(Debug)]
struct Dog;
impl Animal for Dog {
    fn speak(&self) {
        println!("dyn Trait: Woof!");
    }
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn test_dyn_trait() {
        let animal: Box<dyn Animal> = Box::new(Dog);
        // should not panic
        animal.speak();
    }

    #[test]
    fn test_nll() {
        let mut x = 5;
        let y = &x;
        let _val = *y; // last use of y
        x = 6; // NLL: OK because y is no longer used
        assert_eq!(x, 6);
    }

    #[test]
    fn test_async() {
        use futures::executor::block_on;
        async fn add(a: i32, b: i32) -> i32 {
            a + b
        }
        let result = block_on(add(3, 4));
        assert_eq!(result, 7);
    }
}
