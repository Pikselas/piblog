class Hexagon implements HasEnvElement
{
    private static HexagonSide = class 
    {
        private side: EnvElement;
        private sid_type: "top" | "bottom";
        constructor(public type: "top" | "bottom" , width: number = 100, height: number = 100)
        {
            this.sid_type = type;
            const halfWidth = Math.round(width / 2);
            this.side = new EmptyEnvElement();
            this.side.style.width = "0px";
            this.side.style.height = "0px";
            this.side.style.borderStyle = "solid";
            this.side.style.borderWidth = width + "px";
            this.side.style.borderRightWidth = halfWidth + "px";
            this.side.style.borderLeftWidth = halfWidth + "px";
            this.side.style.borderColor = "transparent";
            this.side.style.position = "absolute";
            
            switch (type)
            {
                case "top":
                    this.side.style.borderBottomColor = "black";
                    this.side.style.borderBottomWidth = height + "px";
                    this.side.style.borderTop = "0px";
                    break;
                case "bottom":
                    this.side.style.borderTopColor = "black";
                    this.side.style.borderTopWidth = height + "px";
                    this.side.style.borderBottom = "0px";
                    break;
            }
        }
        public SetColor(color: string)
        {
            switch (this.sid_type)
            {
                case "top":
                    this.side.style.borderBottomColor = color;
                    break;
                case "bottom":
                    this.side.style.borderTopColor = color;
                    break;
            }
        }
        get Element(): EnvElement
        {
            return this.side;
        }
    }
    private hexagon: EnvObject;
    private middle: EnvElement;
    private sideTop: InstanceType<typeof Hexagon.HexagonSide>;
    private sideBottom: InstanceType<typeof Hexagon.HexagonSide>;
    constructor(width: number , height: number)
    {
        this.hexagon = new EmptyEnvObject();
        const oneThirdHeight = Math.round(height / 3);
        this.sideTop = new Hexagon.HexagonSide("top",width, oneThirdHeight);
        this.sideBottom = new Hexagon.HexagonSide("bottom" , width, oneThirdHeight);
        
        this.sideTop.Element.style.top = "0px";
        this.sideTop.Element.style.left = "0px";

        this.sideBottom.Element.style.top = height - oneThirdHeight + "px";
        this.sideBottom.Element.style.left = "0px";
        
        this.middle = new EmptyEnvElement();
        this.middle.style.width = width + "px";
        this.middle.style.height = oneThirdHeight + 3 + "px";
        this.middle.style.position = "absolute";
        this.middle.style.top = oneThirdHeight - 1 + "px";
        this.middle.style.backgroundColor = "black";

        this.hexagon = new EmptyEnvObject();
        this.hexagon.style.width = width + "px";
        this.hexagon.style.height = height + "px";
        this.hexagon.style.position = "relative";
        
        this.hexagon.addElement(this.sideTop.Element);
        this.hexagon.addElement(this.middle);
        this.hexagon.addElement(this.sideBottom.Element);
    }
    public SetColor(color: string)
    {
        this.sideTop.SetColor(color);
        this.sideBottom.SetColor(color);
        this.middle.style.backgroundColor = color;
    }
    public SetStyle(styleName: string,styleValue: string)
    {
        this.hexagon.style[styleName] = styleValue;
    }
    addElement(element:EnvElement):void
    {
        this.hexagon.addElement(element);
    }
    get Elements(): EnvElement[]
    {
        return this.hexagon.Elements;
    }
    get style():CSSStyleDeclaration
    {
        return this.hexagon.style;
    }
    get Element(): EnvElement
    {
        return this.hexagon;
    }
}