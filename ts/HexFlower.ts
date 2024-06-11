
interface HexFlower extends HasEnvElement , HasEnvObject
{
    SetColor(color: string): void;
    SetPetalColor(row: number, petal: number, color: string): void;
    GetPetalCount(row: number): number;
    PetalRows: number;
    BackgroundStyle: CSSStyleDeclaration;
}

class flowerBase implements HexFlower
{
    protected flower: EnvObject;
    protected flowerBackground: EnvElement;
    protected hex_row: HexRow[];

    constructor()
    {
        this.flower = new EmptyEnvObject();
        this.flowerBackground = new EmptyEnvElement();
        this.hex_row = [];
    }

    SetColor(color: string)
    {
        this.hex_row.forEach((row) => row.SetColor(color));
    }
    SetPetalColor(row: number , col: number , color: string)
    {
        this.hex_row[row].SetIndividualColor(col,color);
    }
    GetPetalCount(row: number): number
    {
        return this.hex_row[row].TotalHexes;
    }
    get BackgroundStyle(): CSSStyleDeclaration
    {
        return this.flowerBackground.style;
    }
    get Element(): EnvElement
    {
        return this.flower;
    }
    get Object(): EnvObject
    {
        return this.flower;
    }
    get PetalRows(): number
    {
        return this.hex_row.length;
    }
}

class HexClover extends flowerBase
{

    constructor(width: number, height: number)
    {
        super();
        this.hex_row = [ new HexRow(1, width / 2, height / 2) , new HexRow(2, width, height / 2)];

        this.hex_row[0].Element.style.marginLeft = width / 4 - 1 + "px";
        this.hex_row[1].Element.style.marginTop = -height / 6 - 2 + "px";

        this.flowerBackground.style.position = "absolute";
        this.flowerBackground.style.width = Math.floor(width / 2) - 4 + "px";
        this.flowerBackground.style.height = height / 2 + "px";
        this.flowerBackground.style.marginLeft = Math.floor(width / 4) + height / width + "px";
        this.flowerBackground.style.marginTop = height / 6 + "px";

        this.flower.addElement(this.flowerBackground);
        this.flower.addElement(AsElement(this.hex_row[0]));
        this.flower.addElement(AsElement(this.hex_row[1]));
    }
}

class HexSunflower extends flowerBase
{
    constructor(width: number, height: number) 
    {
        super();
        //First and Last row shares 1/3 of the height with the middle row
        //The middle row shares 2/3 of the height with the first and last row
        
        //The height of the middle row is same as the height of the flower
        //The height of the first and last row is 2/3 as the height of the flower
        
        const mainHeight = height / 3;
        const perRowHeight = mainHeight + (height - mainHeight * (2/3 + 1/3 + 2/3)) / 5;
        const twoThirdWidth = Math.round(width * (2 / 3)) + 1;
        
        this.hex_row.push(new HexRow(2,twoThirdWidth, perRowHeight));
        this.hex_row.push(new HexRow(3,width ,perRowHeight));
        this.hex_row.push(new HexRow(2, twoThirdWidth, perRowHeight));

        //The first and last row is shifted to the left by 1/3 of the width of the flower
       
        AsObject(this.hex_row[0]).style.marginLeft = twoThirdWidth - width / 2 - 3  + "px";
        AsObject(this.hex_row[2]).style.marginLeft = twoThirdWidth - width / 2 - 3  + "px";
       
        //The second and last row is shifted to the top by 1/3 of the height of them
       
        AsObject(this.hex_row[1]).style.marginTop =  -perRowHeight * (1/3) + "px";
        AsObject(this.hex_row[2]).style.marginTop = -perRowHeight * (1/3) + "px";

        this.flowerBackground.style.position = "absolute";
        this.flowerBackground.style.width = twoThirdWidth - 7 + "px";
        this.flowerBackground.style.height = height - (perRowHeight - (perRowHeight * 1 / 3))  + "px";
        this.flowerBackground.style.marginLeft = twoThirdWidth + 2 - width / 2 - 3  + "px";
        this.flowerBackground.style.marginTop = perRowHeight - (perRowHeight * 2 / 3) + "px";

        this.flower.addElement(this.flowerBackground);
        this.flower.addElement(AsElement(this.hex_row[0]));
        this.flower.addElement(AsElement(this.hex_row[1]));
        this.flower.addElement(AsElement(this.hex_row[2]));
        this.flower.style.width = width + "px";
    }
}