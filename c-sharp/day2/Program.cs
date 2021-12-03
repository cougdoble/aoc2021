// not proud of this, but it is what it is :(
static int CalcDepthAndPosition(string filePath)
{
	var input = new List<string>();
	using (var streamReader = new StreamReader(new FileStream(filePath, FileMode.Open, FileAccess.Read)))
	{
		string? line;
		while ((line = streamReader.ReadLine()) is not null)
		{
			input.Add(line);
		}
	}

	var horizontalPosition = 0;
	var depth = 0;

	foreach (var movement in input.Select(i => i.Split(" ")))
	{
		var amount = int.Parse(movement[1]);
		switch (movement[0])
		{
			case "forward":
				horizontalPosition += amount;
				break;
			case "down":
				depth += amount;
				break;
			case "up":
				depth -= amount;
				break;
		}
	}

	return horizontalPosition * depth;
}

static int CalcDepthAndPosition2(string filePath)
{
	var input = new List<string>();
	using (var streamReader = new StreamReader(new FileStream(filePath, FileMode.Open, FileAccess.Read)))
	{
		string? line;
		while ((line = streamReader.ReadLine()) is not null)
		{
			input.Add(line);
		}
	}

	var horizontalPosition = 0;
	var depth = 0;
	var aim = 0;

	foreach (var movement in input.Select(i => i.Split(" ")))
	{
		var amount = int.Parse(movement[1]);
		switch (movement[0])
		{
			case "forward":
				horizontalPosition += amount;
				depth += amount * aim;
				break;
			case "down":
				aim += amount;
				break;
			case "up":
				aim -= amount;
				break;
		}
	}

	Console.WriteLine($"horizontal position: {horizontalPosition}");
	Console.WriteLine($"depth: {depth}");
	return horizontalPosition * depth;
}

const string testFilePath = @"day2-test-input.txt";
const string filePath = @"day2-input.txt";

Console.WriteLine(CalcDepthAndPosition(filePath));
Console.WriteLine(CalcDepthAndPosition2(filePath));