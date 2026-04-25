fn main() -> Result<(), Box<dyn std::error::Error>> {
    println!("--- Rust 2015 (Core Features) Demo ---");

    // 1. Ownership & Borrowing
    // 所有権の移動（Move）と不変借用（Borrow）の基本
    let s = String::from("hello");
    print_string(&s); // 借用: 所有権を貸し出す
    let _s2 = s;     // 移動: 所有権が _s2 へ移る
    // println!("{}", s); // コンパイルエラー: 所有権が移動済みのため

    // 2. Pattern Matching
    // Option型とパターンマッチによる安全な値の取り出し
    let x = Some(5);
    match x {
        Some(i) => println!("Pattern Matching: got Some({})", i),
        None => println!("Pattern Matching: got None"),
    }

    // 3. Trait & Struct
    // 抽象化のためのトレイトとその実装
    let circle = Circle { radius: 1.0 };
    println!("Trait: Area of circle is {:.3}", circle.area());
    println!("Debug Print: {:?}", circle);

    Ok(())
}

fn print_string(s: &str) {
    println!("Ownership & Borrowing: {}", s);
}

trait HasArea {
    fn area(&self) -> f64;
}

#[derive(Debug)]
struct Circle {
    radius: f64,
}

impl HasArea for Circle {
    fn area(&self) -> f64 {
        std::f64::consts::PI * self.radius * self.radius
    }
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn test_circle_area() {
        let c = Circle { radius: 2.0 };
        // 浮動小数点の比較は近似値で行う
        assert!((c.area() - 12.566).abs() < 0.01);
    }
}
