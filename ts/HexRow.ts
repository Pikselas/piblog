class HexRow implements HasEnvElement , HasEnvObject
{
    private row: EnvObject;
    private hexes: Hexagon[];
    constructor(count: number , width: number , height: number)
    {
        this.row = new EmptyEnvObject();
        this.hexes = [];
        const perHexWidth = Math.round((width / count) - count - 2);
        for (let i = 0; i < count; i++)
        {
            let hex = new Hexagon(perHexWidth,height);
            this.hexes.push(hex);
            this.row.addElement(AsElement(hex));
            hex.SetStyle("display","inline-block");
            hex.SetStyle("margin-left","2px");
        }
        this.row.style.width = width + "px";
    }
    SetColor(color: string)
    {
        this.hexes.forEach((hex) => hex.SetColor(color));
    }
    SetIndividualColor(index: number , color: string)
    {
        this.hexes[index].SetColor(color);
    }
    get Element(): EnvElement
    {
        return this.row;
    }
    get Object(): EnvObject
    {
        return this.row;
    }
    get TotalHexes(): number
    {
        return this.hexes.length;
    }
}