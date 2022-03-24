## How to run it

Download this repository or the single executable file `symmetric_group` (without suffix), then enter the project folder, run script:

Non-recursive version:
```bash
./symmetric_group <n>
```
Non-recursive version with multiple threads:
```bash
./symmetric_group <n> -m
```
Or the recursive version:
```bash
./symmetric_group <n> -r
```

The character table will then be saved to the current directory with the name like `output/character_table (S_5).txt`.

### Example:
```
>> ./symmetric_group 5 -r
[Character table of Sysmmetric Group (n=5)] Recursive Version
Partition Length: 7
Calculating...
Time Spent: 0 ms
Output file: character_table (S_5).txt
Done!
```

Content of the file `character_table (S_5).txt`:
```
Character table of Sysmmetric Group (n=5)
7X7	[1 1 1 1 1]	[2 1 1 1]	[2 2 1]	[3 1 1]	[3 2]	[4 1]	[5]	
[1 1 1 1 1]	1	-1	1	1	-1	-1	1	
[2 1 1 1]	4	-2	0	1	1	0	-1	
[2 2 1]	5	-1	1	-1	-1	1	0	
[3 1 1]	6	0	-2	0	0	0	1	
[3 2]	5	1	1	-1	1	-1	0	
[4 1]	4	2	0	1	-1	0	-1	
[5]	1	1	1	1	1	1	1	
```

## Test Results

For n=10,15,20,25,30,35, we have test results (not include file write time):

|n|partition size|time spend (non-recursive version, ms)|time spend (non-recursive version,multiple thread, ms)| time spend (recursive version, ms)|
|-|--------------|---------------|-----------------|-|
|5|7|0|0|0|
|10|42|1|0|1|
|15|176|17|4|21|
|20|627|151|46|270|
|25|1958|1524|434|2886|
|30|5604|12254|3294|24189|
|35|14883|92343|25044|216933|

Test device: MacBook Pro 13, 2.4 GHz Quad-Core Intel Core i5, 8 GB 2133 MHz LPDDR3

## Download

The executable file: [symmetric_group](https://github.com/youxingz/symmetric_group_character/blob/master/symmetric_group)

Output (n=2,...,34) files: [outputs](https://github.com/youxingz/symmetric_group_character/tree/master/output)

Because of the size of output files, we compress these files from `1500M` to `374M` with 7z.
