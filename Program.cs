using System.Text;

static int DepthSum(string filePath)
{
    var depthIncreasedCount = 0;
    int? previousValue = null;

    foreach (var line in File.ReadLines(filePath, Encoding.UTF8))
    {
        var currentLine = int.Parse(line);
        previousValue ??= currentLine;

        if (currentLine > previousValue)
        {
            depthIncreasedCount++;
        }

        previousValue = currentLine;
    }

    return depthIncreasedCount;
}

static int MaxDepthSum(string filePath)
{
    // get values from file
    var depthData = new List<int>();
    using (var streamReader = new StreamReader(new FileStream(filePath, FileMode.Open, FileAccess.Read)))
    {
        string? line;
        while ((line = streamReader.ReadLine()) is not null) depthData.Add(int.Parse(line));
    }

    const int windowSize = 3;

    // keep a separate reference that we can mutate at the end of the for loop by calling Skip(1)
    var remaining = depthData.AsEnumerable();

    var rounds = depthData.Count - windowSize;
    var count = 0;

    for (var i = 0; i < rounds; i++)
    {
        // at this starting position in the list of numbers, keep a copy on the right that
        // is advanced by one element
        var left = remaining as int[] ?? remaining.ToArray();
        var right = left.Skip(1);

        // take the chosen number of elements (1 for part one, 3 for part two) and sum them
        var leftSum = left.Take(windowSize).Sum();
        var rightSum = right.Take(windowSize).Sum();

        if (leftSum < rightSum)
        {
            count++;
        }

        // advance to the next starting position (this is done rounds times)
        remaining = left.Skip(1);
    }

    return count;
}

Console.WriteLine($"Part 1: {DepthSum(@"c:\Users\dcoble\Desktop\aoc1-input.txt")}");
Console.WriteLine($"Part 2: {MaxDepthSum(@"c:\Users\dcoble\Desktop\aoc1-input.txt")}");