use std::collections::HashMap;

pub fn p1(input: Vec<isize>, days: isize) -> usize {
    println!("Initial state: {:?}", input);

    let mut fishes = vec![0, 0, 0, 0, 0, 0, 0, 0, 0];
    for fish in input {
        fishes[fish as usize] += 1
    }
    println!("{:?}", fishes);

    for i in 0..days {
        // println!("After {} days: {:?}", i + 1, fishes);
        println!("day {} has {} fishes", i, fishes.iter().sum::<isize>());

        let mut totals_day: Vec<isize> = vec![0, 0, 0, 0, 0, 0, 0, 0, 0];
        for (days, num_fish) in fishes.iter_mut().enumerate() {
            if days == 0 {
                totals_day[8] += *num_fish;
                totals_day[6] += *num_fish;
                totals_day[0] -= *num_fish;
                continue;
            }
            totals_day[days] -= *num_fish;
            totals_day[(days - 1)] += *num_fish;
        }

        for (days, num_fish) in totals_day.iter().enumerate() {
            fishes[days] += num_fish;
        }
    }

    fishes.iter().sum::<isize>() as usize
}

#[cfg(test)]
mod tests {
    use crate::*;

    #[test]
    fn test_p1() {
        let input = vec![3, 4, 3, 1, 2];
        assert_eq!(p1(input, 80), 5934);
    }

    #[test]
    fn test_p2() {
        let input = vec![3, 4, 3, 1, 2];
        assert_eq!(p1(input, 256), 26984457539);
    }

    #[test]
    fn test_input() {
        let input = vec![
            1, 2, 4, 5, 5, 5, 2, 1, 3, 1, 4, 3, 2, 1, 5, 5, 1, 2, 3, 4, 4, 1, 2, 3, 2, 1, 4, 4, 1,
            5, 5, 1, 3, 4, 4, 4, 1, 2, 2, 5, 1, 5, 5, 3, 2, 3, 1, 1, 3, 5, 1, 1, 2, 4, 2, 3, 1, 1,
            2, 1, 3, 1, 2, 1, 1, 2, 1, 2, 2, 1, 1, 1, 1, 5, 4, 5, 2, 1, 3, 2, 4, 1, 1, 3, 4, 1, 4,
            1, 5, 1, 4, 1, 5, 3, 2, 3, 2, 2, 4, 4, 3, 3, 4, 3, 4, 4, 3, 4, 5, 1, 2, 5, 2, 1, 5, 5,
            1, 3, 4, 2, 2, 4, 2, 2, 1, 3, 2, 5, 5, 1, 3, 3, 4, 3, 5, 3, 5, 5, 4, 5, 1, 1, 4, 1, 4,
            5, 1, 1, 1, 4, 1, 1, 4, 2, 1, 4, 1, 3, 4, 4, 3, 1, 2, 2, 4, 3, 3, 2, 2, 2, 3, 5, 5, 2,
            3, 1, 5, 1, 1, 1, 1, 3, 1, 4, 1, 4, 1, 2, 5, 3, 2, 4, 4, 1, 3, 1, 1, 1, 3, 4, 4, 1, 1,
            2, 1, 4, 3, 4, 2, 2, 3, 2, 4, 3, 1, 5, 1, 3, 1, 4, 5, 5, 3, 5, 1, 3, 5, 5, 4, 2, 3, 2,
            4, 1, 3, 2, 2, 2, 1, 3, 4, 2, 5, 2, 5, 3, 5, 5, 1, 1, 1, 2, 2, 3, 1, 4, 4, 4, 5, 4, 5,
            5, 1, 4, 5, 5, 4, 1, 1, 5, 3, 3, 1, 4, 1, 3, 1, 1, 4, 1, 5, 2, 3, 2, 3, 1, 2, 2, 2, 1,
            1, 5, 1, 4, 5, 2, 4, 2, 2, 3,
        ];
        assert_eq!(p1(input, 256), 352872);
    }
}
