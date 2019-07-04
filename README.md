# trackermod

Library for Amiga Protracker .mod module music tracker format for Go. 

## Format

Amiga Protracker .mod module format. [Endianness](https://en.wikipedia.org/wiki/Endianness) is big.


 Offset  | Length | Type    | Description 
---------|--------|---------|---
  0      | 22     | byte    | Song title       

Sound sample data (30 bytes per sample)
 
 Offset  | Length | Type    | Description 
---------|--------|---------|---
 0 (20)  | 22     | byte    | Sample name 
 22 (42) | 2      | uint16  | Length as number of uint16's (multiply with 2 for real byte count)
 24 (44) | 1      | uint8   | Finetune with mask 0b00001111
 25 (45) | 1      | uint8   | Volume
 26 (46) | 2      | uint16  | Repeat point as number of uint16's (multiply with 2 for real byte count)
 28 (48) | 2      | uint16  | Repeat length as number of uint16's (multiply with 2 for real byte count)

Next samples:

 Offset  | Length | Type    | Description 
---------|--------|---------|---
 50      | 30     | byte    | Sample #2 
 80      | 30     | byte    | Sample #3 
 ...     | ...    | ...     | ... 
 890     | 30     | byte    | Sample #30 
 920     | 30     | byte    | Sample #31 

Pattern info:

 Offset  | Length | Type    | Description 
---------|--------|---------|---
 950     | 1      | uint8   | Song length in patterns 
 951     | 1      | uint8   | Restart. Not used. Set to 127 in old trackers
 952     | 128    | uint8   | Pattern positions 0-127. Value can be 0-63.
 1080    | 4      | byte    | Magic "M.K." (stands for Michael Kleps)

### Pattern data 

Get the number of patterns from offset 950 mentioned earlier and loop that many times. Maximum is 64.

Pattern information is stored as

 Row  | Channel 1 | Channel 2 | Channel 3 | Channel 4  
------|-----------|-----------|-----------|---
 00   | ch 1 note | ch 2 note | ch 3 note | ch 4 note
 01   | ch 1 note | ch 2 note | ch 3 note | ch 4 note
 ...  | ...       | ...       | ...       | ... 
 62   | ch 1 note | ch 2 note | ch 3 note | ch 4 note
 63   | ch 1 note | ch 2 note | ch 3 note | ch 4 note 

and then next pattern starts with same way.


 Offset      | Length | Type    | Description 
-------------|--------|---------|---
 0 (1084)    | 4      | uint32  | Note data for channel #1 pattern 0 index 0
 4 (1088)    | 4      | uint32  | Note data for channel #2 pattern 0 index 0
 8 (1092)    | 4      | uint32  | Note data for channel #3 pattern 0 index 0
 12 (1096)   | 4      | uint32  | Note data for channel #4 pattern 0 index 0
 16 (1100)   | 4      | uint32  | Note data for channel #1 pattern 0 index 1
 10 (1088)   | 4      | uint32  | Note data for channel #2 pattern 0 index 1
 10 (1092)   | 4      | uint32  | Note data for channel #3 pattern 0 index 1
 10 (1096)   | 4      | uint32  | Note data for channel #4 pattern 0 index 1
 ...         | ...    | ...     | ...
 10 (1096)   | 4      | uint32  | Note data for channel #1 pattern 0 index 63
 10 (1096)   | 4      | uint32  | Note data for channel #2 pattern 0 index 63
 10 (1096)   | 4      | uint32  | Note data for channel #3 pattern 0 index 63
 1020 (1096) | 4      | uint32  | Note data for channel #4 pattern 0 index 63


 Offset | Length | Type    | Description 
--------|--------|---------|---
 2108   | 1024   | uint32  | Pattern 1
 3132   | 1024   | uint32  | Pattern 2
 ...    | ...    | ...     | ...
 64572  | 1024   | uint32  | Pattern 62 
 65596  | 1024   | uint32  | Pattern 63 (max)


### Note data (uint32)

 Bit mask                         | Hex        | Bits | Description 
----------------------------------|------------|------|---
 11110000000000000000000000000000 | 0xF0000000 | 4    | Sample's bits upper
 00001111111111110000000000000000 | 0x0FFF0000 | 12   | Period
 00000000000000001111000000000000 | 0x0000F000 | 4    | Sample's bits lower
 00000000000000000000111100000000 | 0x00000F00 | 4    | Effect 0x0 - 0xF
 00000000000000000000000011111111 | 0x000000FF | 8    | Effect parameters

To get correct sample number you must combine sample's upper and lower bits to get the uint8 needed. 


### Sound wave data

...

### Notes

 Note | Binary           | Binary reversed  | Hex    | Dec
------|------------------|------------------|--------|----
 A#1  | 0000000111100000 | 0000011110000000 | 0x01e0 | 480
 A#2  | 0000000011110000 | 0000111100000000 | 0x00f0 | 240
 A#3  | 0000000001111000 | 0001111000000000 | 0x0078 | 120
 A-1  | 0000000111111100 | 0011111110000000 | 0x01fc | 508
 A-2  | 0000000011111110 | 0111111100000000 | 0x00fe | 254
 A-3  | 0000000001111111 | 1111111000000000 | 0x007f | 127
 B-1  | 0000000111000101 | 1010001110000000 | 0x01c5 | 453
 B-2  | 0000000011100010 | 0100011100000000 | 0x00e2 | 226
 B-3  | 0000000001110001 | 1000111000000000 | 0x0071 | 113
 C#1  | 0000001100101000 | 0001010011000000 | 0x0328 | 808
 C#2  | 0000000110010100 | 0010100110000000 | 0x0194 | 404
 C#3  | 0000000011001010 | 0101001100000000 | 0x00ca | 202
 C-1  | 0000001101011000 | 0001101011000000 | 0x0358 | 856
 C-2  | 0000000110101100 | 0011010110000000 | 0x01ac | 428
 C-3  | 0000000011010110 | 0110101100000000 | 0x00d6 | 214
 D#1  | 0000001011010000 | 0000101101000000 | 0x02d0 | 720
 D#2  | 0000000101101000 | 0001011010000000 | 0x0168 | 360
 D#3  | 0000000010110100 | 0010110100000000 | 0x00b4 | 180
 D-1  | 0000001011111010 | 0101111101000000 | 0x02fa | 762
 D-2  | 0000000101111101 | 1011111010000000 | 0x017d | 381
 D-3  | 0000000010111110 | 0111110100000000 | 0x00be | 190
 E-1  | 0000001010100110 | 0110010101000000 | 0x02a6 | 678
 E-2  | 0000000101010011 | 1100101010000000 | 0x0153 | 339
 E-3  | 0000000010101010 | 0101010100000000 | 0x00aa | 170
 F#1  | 0000001001011100 | 0011101001000000 | 0x025c | 604
 F#2  | 0000000100101110 | 0111010010000000 | 0x012e | 302
 F#3  | 0000000010010111 | 1110100100000000 | 0x0097 | 151
 F-1  | 0000001010000000 | 0000000101000000 | 0x0280 | 640
 F-2  | 0000000101000000 | 0000001010000000 | 0x0140 | 320
 F-3  | 0000000010100000 | 0000010100000000 | 0x00a0 | 160
 G#1  | 0000001000011010 | 0101100001000000 | 0x021a | 538
 G#2  | 0000000100001101 | 1011000010000000 | 0x010d | 269
 G#3  | 0000000010000111 | 1110000100000000 | 0x0087 | 135
 G-1  | 0000001000111010 | 0101110001000000 | 0x023a | 570
 G-2  | 0000000100011101 | 1011100010000000 | 0x011d | 285
 G-3  | 0000000010001111 | 1111000100000000 | 0x008f | 143