// style-rule : js-style-property

// rules for the flower background
const FlowerBackgroundRules = { "background-color" : "backgroundColor" , "background" : "background" , "animation" : "animation" }

function SetFlowerTheme(flower: HexFlower , className: string)
{
    // get the necessary style rules
    var dump_obj = document.createElement('div');
    dump_obj.className = className;
    dump_obj.hidden = true;
    document.body.appendChild(dump_obj);
    var style = window.getComputedStyle(dump_obj);

    // retrieve the rules for the flower background
    Object.keys(FlowerBackgroundRules).forEach((rule) => 
    {
        var prop = FlowerBackgroundRules[rule];
        var val = style.getPropertyValue(rule);
        flower.BackgroundStyle[prop] = val;
    });

    // retrieve the rules for the flower petal
    flower.SetColor(style.getPropertyValue("color"));
    for(let i = 0 ; i < flower.PetalRows ; ++i)
    {
        for(let j = 0 ; j < flower.GetPetalCount(i) ; ++j)
        {
            const petalColor = style.getPropertyValue(`--petal${i}${j}`);
            if (petalColor != "")
            {
                flower.SetPetalColor(i,j,petalColor);
            }
        }
    }
    dump_obj.remove();
}