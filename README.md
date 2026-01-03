# PXM (PGM XML Merger)

PXM is a CLI tool for merging **[PGM](https://github.com/PGMDev/PGM)** XML files into a single output.

## Usage

### Command Format

`./pxm --folder <folder> --base <base.xml> --out <output.xml>`

### Flags

- `--folder <folder>`: Folder with XML files to merge. Default: src
- `--base <base.xml>`: Base XML file. Default: base.xml
- `--out <output.xml>`: Output file. Default: map.xml

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