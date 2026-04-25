fn main() -> Result<(), Box<dyn std::error::Error>> {
    println!("--- Rust 2024 Demo ---");

    // 1. RPITIT (ReturnTypeInsideTrait) & impl Trait capture rules
    // 2024 では impl Trait がデフォルトですべてのインスコープなライフタイムをキャプチャする
    // これにより、以前の editoin で必要だった「+ 'a」などの指定が簡略化される
    let message = String::from("Rust 2024 captures lifetimes by default");
    let printer = create_printer(&message);
    printer.print();

    // 2. Future expansion: gen blocks (Draft / Preview)
    // 注意: 現在の安定版や初期の2024では、genキーワードは予約語としての側面が強い
    // 将来的にはネイティブなジェネレータ（Iteratorを生成する構文）として機能する予定
    println!("Future: 'gen' is now a reserved keyword for native generators.");

    Ok(())
}

trait Printer {
    fn print(&self);
}

struct SimplePrinter<'a> {
    msg: &'a str,
}

impl<'a> Printer for SimplePrinter<'a> {
    fn print(&self) {
        println!("impl Trait: {}", self.msg);
    }
}

// Rust 2024 では -> impl Printer で 'a が暗黙的にキャプチャされる
fn create_printer<'a>(msg: &'a str) -> impl Printer {
    SimplePrinter { msg }
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn test_capture_rules() {
        let msg = String::from("test");
        let printer = create_printer(&msg);
        printer.print();
    }
}
