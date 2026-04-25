fn main() -> Result<(), Box<dyn std::error::Error>> {
    println!("--- Rust 2015 (Core Features) Learning Demo ---");

    // 1. Ownership & Borrowing (所有権と借用)
    // ---------------------------------------------------------
    // Rustの最も重要な概念。メモリの安全性を保証します。
    let s1 = String::from("hello");
    let s2 = s1; // 所有権の移動 (Move): s1はこれ以降使えません
    // println!("{}", s1); // コンパイルエラー！

    let s3 = String::from("world");
    print_string(&s3); // 借用 (Borrowing): 所有権を貸し出し、読み取り専用でアクセス
    println!("Still have s3: {}", s3); // 貸しただけなので、まだ使えます

    let mut s4 = String::from("editable");
    modify_string(&mut s4); // 可変の借用 (Mutable Borrowing): データを変更可能
    println!("Modified s4: {}", s4);

    // 2. Pattern Matching & Enums (列挙型とパターンマッチ)
    // ---------------------------------------------------------
    // 値の構造に基づいて分岐を行います。
    let msg = Message::Write(String::from("Rust is safe"));
    match msg {
        Message::Quit => println!("Quit variant"),
        Message::Move { x, y } => println!("Move to x:{}, y:{}", x, y),
        Message::Write(text) => println!("Message: {}", text),
        Message::ChangeColor(r, g, b) => println!("Color: R{}, G{}, B{}", r, g, b),
    }

    // 3. Option & Result (エラーハンドリング)
    // ---------------------------------------------------------
    // nullの代わりに Option、例外の代わりに Result を使います。
    let divided = divide(10.0, 2.0);
    match divided {
        Ok(val) => println!("Division result: {}", val),
        Err(e) => println!("Error: {}", e),
    }

    // 4. Trait & Generics (トレイトとジェネリクス)
    // ---------------------------------------------------------
    // 多態性（ポリモーフィズム）を実現します。
    let circle = Circle { radius: 1.0 };
    print_area(circle);

    let rectangle = Rectangle { width: 4.0, height: 2.0 };
    print_area(rectangle);

    // 5. Lifetimes (ライフタイム)
    // ---------------------------------------------------------
    // 参照が「いつまで有効か」をコンパイラに伝えます（基本は自動ですが、明示が必要な場合もあります）。
    let string1 = String::from("long string");
    let string2 = "xyz";
    let result = longest(string1.as_str(), string2);
    println!("The longest string is {}", result);

    Ok(())
}

// --- Supporting Code ---

fn print_string(s: &str) {
    println!("Borrowed: {}", s);
}

fn modify_string(s: &mut String) {
    s.push_str(" content");
}

enum Message {
    Quit,
    Move { x: i32, y: i32 },
    Write(String),
    ChangeColor(i32, i32, i32),
}

fn divide(numerator: f64, denominator: f64) -> Result<f64, String> {
    if denominator == 0.0 {
        // 動的なエラーメッセージの生成（&'static str では不可）
        Err(format!("Cannot divide {} by zero", numerator))
    } else {
        Ok(numerator / denominator)
    }
}

trait HasArea {
    fn area(&self) -> f64;
}

#[derive(Debug)]
struct Circle { radius: f64 }
impl HasArea for Circle {
    fn area(&self) -> f64 { std::f64::consts::PI * self.radius * self.radius }
}

#[derive(Debug)]
struct Rectangle { width: f64, height: f64 }
impl HasArea for Rectangle {
    fn area(&self) -> f64 { self.width * self.height }
}

// ジェネリクスを用いた関数
fn print_area<T: HasArea>(shape: T) {
    println!("The area is {:.2}", shape.area());
}

// ライフタイム注釈: 'a は、返される参照が引数の参照と同じ期間だけ有効であることを保証します
fn longest<'a>(x: &'a str, y: &'a str) -> &'a str {
    if x.len() > y.len() { x } else { y }
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn test_circle_area() {
        let c = Circle { radius: 2.0 };
        assert!((c.area() - 12.566).abs() < 0.01);
    }

    #[test]
    fn test_divide() {
        let result = divide(10.0, 0.0);
        assert!(result.is_err());
        assert_eq!(result.unwrap_err(), "Cannot divide 10 by zero");
        
        assert_eq!(divide(10.0, 2.0).unwrap(), 5.0);
    }
}
