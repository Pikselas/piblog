interface EnvElement
{
    get HtmlElement():HTMLElement;
    get style():CSSStyleDeclaration;
}

interface EnvObject extends EnvElement
{
    addElement(element:EnvElement):void;
    get Elements(): EnvElement[];
}

interface HasEnvElement
{
    get Element():EnvElement;
}

interface HasEnvObject
{
    get Object():EnvObject;
}

class EmptyEnvElement
{ 
    private HTMLElement: HTMLDivElement;
    constructor()
    {
        this.HTMLElement = document.createElement("div");
    }
    get HtmlElement():HTMLElement
    {
        return this.HTMLElement;
    }
    get style():CSSStyleDeclaration
    {
        return this.HTMLElement.style;
    }
}

class EmptyEnvObject extends EmptyEnvElement implements EnvElement, EnvObject
{
    private element: EnvElement[];
    constructor()
    {
        super();
        this.element = [];
    }
    addElement(element: EnvElement): void
    {
        this.HtmlElement.appendChild(element.HtmlElement);
        this.element.push(element);
    }
    get Elements(): EnvElement[]
    {
        return this.element;
    }
}

function AsElement(element:HasEnvElement):EnvElement
{
    return element.Element;
}

function AsObject(element:HasEnvObject):EnvObject
{
    return element.Object;
}
