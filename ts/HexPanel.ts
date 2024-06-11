class HexPanel implements HasEnvElement , HasEnvObject
{
    private panel: EnvObject;
    private rows: HexRow[];
    private panelBackground: EnvElement;
    constructor(rowCount: number , hexCount: number , width: number , height: number)
    {
        this.panel = new EmptyEnvObject();
        this.panelBackground = new EmptyEnvElement();
        this.panel.addElement(this.panelBackground);
        this.panelBackground.style.width = width + "px";
        this.panelBackground.style.height = height + "px";
        this.panelBackground.style.position = "absolute";
        this.panel.style.width = width + "px";
        this.panel.style.height = height + "px";
        this.panel.style.overflow = "hidden";
        this.rows = [];
        const mainHeight = height / rowCount;
        const perHexHeight = Math.round(mainHeight + (height - mainHeight * (2/3 + (rowCount - 2) * 1/3 + 2/3)) / rowCount - 3);
        let isEvenRow = true;
        const perHexWidth = Math.round((width / hexCount));
        console.log(hexCount , perHexWidth , width , perHexWidth * hexCount);
        for (let i = 0; i < rowCount + 1; i++)
        {
            let row = new HexRow(hexCount + 2 ,width + perHexWidth * 2,perHexHeight);
            this.rows.push(row);
            this.panel.addElement(AsElement(row));
            AsElement(row).style.marginTop = -Math.round(perHexHeight / 3) + "px";
            isEvenRow ? AsElement(row).style.marginLeft =  null : AsElement(row).style.marginLeft = -Math.round(perHexWidth / 2) + 4 + "px";
            isEvenRow = !isEvenRow;
        }
        this.panel.style.height = height + "px";
    }
    SetColor(color: string)
    {
        this.rows.forEach((row) => row.SetColor(color));
    }
    get BackgroundStyle(): CSSStyleDeclaration
    {
        return this.panelBackground.style;
    }
    get Element(): EnvElement
    {
        return this.panel;
    }
    get Object(): EnvObject
    {
        return this.panel;
    }
}