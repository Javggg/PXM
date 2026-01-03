# PXM (PGM XML Merger)

PXM is a CLI tool for merging **[PGM](https://github.com/pgm)** XML files into a single output.

## Usage

### Command Format

./pxm --folder <folder> --base <base.xml> --out <output.xml>

### Flags

- --folder <folder>: Folder with XML files to merge. Default: src
- --base <base.xml>: Base XML file. Default: base.xml
- --out <output.xml>: Output file. Default: map.xml

### Example

1. Default settings:

./pxm

2. Custom folder, base, and output:

./pxm --folder myxmls --base custom_base.xml --out merged_output.xml

### Modules left to map:
- Falling blocks
- Lootables
- Shops
- Item mods
- Crafting
- Broadcasts
- Rules
- Stats
- Payloads
- Proximity alarms
- Lanes
- Gamerules
- Compass
- Classes
- Pickups
- Blocks drops
- Enderchest

#### Others (base exclusive)
- Score
- Time
- Rage
- Damage
- TNT