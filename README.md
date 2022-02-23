# Clicker

Clicker is a really simple piece of code that aims at clicking at the screen for you at random time.

## How to configure it

In order to use the clicker, you need to fill a config.json file.

This file has optional and mandatory fields.

### Mandatory fields

**position**

Position is an array of mouse positions in pixel (x, y) based on your screen resolution. The more you have values in this array, the more diversified will be your clicker.

**clickType**

Can be `left` or `right` depending on the action you want to do once the mouse has moved


**frequency**

This option will determine at which range of frequency the clicker will click in **seconds** (defined by a min frequency and a max frequency)


### Optional fields

**smooth**

Smooth will tell the cursor to move smoothly from point A to point B

**debug**

Debug will provide additional informations, like writing a log when the mouse clicks for instance


## Example

```json
{
    "positions": [
        {
            "x": 885,
            "y": 495
        },
        {
            "x": 1125,
            "y": 600
        }
    ],
    "clickType": "right",
    "smooth": true,
    "frequency": {
        "min": 10,
        "max": 25
    },
    "debug": true
}
```

Enjoy