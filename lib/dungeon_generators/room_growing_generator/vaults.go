package roomgrowinggenerator

// Outside vaults SHOULD have at least one outer door.
var outsideVaults = [][]string{
	{
		"##########",
		"#........+",
		"#.########",
		"#........#",
		"########.#",
		"#........#",
		"##########",
	},
	{
		"#####    ",
		"#...#    ",
		"#...#    ",
		"#...#####",
		"#.......#",
		"#.......+",
		"#########",
	},
	{
		"  #####  ",
		" ##...## ",
		"##.....##",
		"#.......+",
		"##.....##",
		" ##...## ",
		"  #####  ",
	},
	{
		"   ####   ",
		"   #..#   ",
		"   #..#   ",
		"####..####",
		"#........+",
		"#........+",
		"####..####",
		"   #..#   ",
		"   #..#   ",
		"   ####   ",
	},
}

// Inside vaults will be surrounded by floor tiles.
var insideVaults = [][]string{
	{
		".#.",
		"###",
		".#.",
	},
	{
		"######",
		"#....#",
		"#.##.#",
		"#.##.+",
		"#....#",
		"######",
	},
	{
		"######",
		"#..#..",
		"#..#..",
		"#.....",
		"#..#..",
		"#..#..",
		"######",
	},
}