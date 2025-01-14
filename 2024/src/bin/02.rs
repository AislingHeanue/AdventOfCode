advent_of_code::solution!(2);
use itertools::Itertools;
pub fn part_one(input: &str) -> Option<u32> {
    Some(
        input
            .lines()
            .map(|line| {
                line.split(" ")
                    .map(|num| num.parse::<i32>().unwrap())
                    .tuple_windows()
                    .map(|(first, second)| first - second)
            })
            .filter(|differences| {
                let all_positive = differences.clone().all(|val| val > 0);
                let all_negative = differences.clone().all(|val| val < 0);

                let all_between_one_and_three = differences
                    .clone()
                    .all(|val| val.abs() <= 3 && val.abs() >= 1);

                (all_positive || all_negative) && all_between_one_and_three
            })
            .count()
            .try_into()
            .unwrap(),
    )
}

pub fn part_two(input: &str) -> Option<u32> {
    Some(
        input
            .lines()
            .map(|line| line.split(" ").map(|num| num.parse::<i32>().unwrap()))
            .filter(|nums| {
                let num_list: Vec<i32> = nums.clone().collect();
                (0..num_list.len()).any(|i| {
                    let mut smaller_num_list = num_list.clone();
                    smaller_num_list.remove(i);
                    let difference = smaller_num_list
                        .into_iter()
                        .tuple_windows()
                        .map(|(first, second)| first - second);

                    let all_positive = difference.clone().all(|val| val > 0);
                    let all_negative = difference.clone().all(|val| val < 0);

                    let all_between_one_and_three = difference
                        .clone()
                        .all(|val| val.abs() <= 3 && val.abs() >= 1);

                    (all_positive || all_negative) && all_between_one_and_three
                })
            })
            .count()
            .try_into()
            .unwrap(),
    )
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn test_part_one() {
        let result = part_one(&advent_of_code::template::read_file("examples", DAY));
        assert_eq!(result, Some(2));
    }

    #[test]
    fn test_part_two() {
        let result = part_two(&advent_of_code::template::read_file("examples", DAY));
        assert_eq!(result, Some(5));
    }
}
