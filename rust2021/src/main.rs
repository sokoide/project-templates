fn main() -> Result<(), Box<dyn std::error::Error>> {
    println!("--- Rust 2021 Demo ---");

    // 1. Disjoint capture in closures
    // クロージャが構造体全体ではなくフィールド単位で借用を行う
    struct Player {
        name: String,
        score: u32,
    }
    let mut p = Player {
        name: String::from("Alice"),
        score: 100,
    };

    let c = || {
        println!("Closure: Player name is {}", p.name); // p.name を不変借用
    };

    // Rust 2018 までは p 全体が借用されていたため、以下の score 変更はエラーだった
    p.score += 10; // p.score を可変借用（2021 なら OK）

    c();
    println!("Disjoint capture: score updated to {}", p.score);

    // 2. IntoIterator for arrays
    // 配列を直接値として反復できる（.iter() 不要）
    let arr = [1, 2, 3];
    println!("IntoIterator for arrays:");
    for x in arr {
        // 2018 では &arr または arr.iter() が必要だった
        print!("{} ", x);
    }
    println!();

    // 3. New Prelude (TryInto)
    // TryInto が Prelude に入ったため、個別の use なしで型変換が可能
    let n: i64 = 1024;
    let n_i32: i32 = n.try_into()?;
    println!("New Prelude: converted i64 to i32: {}", n_i32);

    Ok(())
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn test_disjoint_capture() {
        struct Player {
            name: String,
            score: u32,
        }
        let mut p = Player {
            name: String::from("Alice"),
            score: 100,
        };
        let c = || {
            let _ = &p.name; // capture only p.name
        };
        p.score += 10; // disjoint: p.score is not captured
        c();
        assert_eq!(p.score, 110);
    }

    #[test]
    fn test_into_iterator_for_arrays() {
        let arr = [1, 2, 3];
        let mut sum = 0;
        for x in arr {
            sum += x;
        }
        assert_eq!(sum, 6);
    }

    #[test]
    fn test_try_into() -> Result<(), Box<dyn std::error::Error>> {
        let n: i64 = 1024;
        let n_i32: i32 = n.try_into()?;
        assert_eq!(n_i32, 1024);
        Ok(())
    }

    #[test]
    fn test_try_into_overflow() {
        let n: i64 = 9999999999999;
        let result: Result<i32, _> = n.try_into();
        assert!(result.is_err());
    }
}
