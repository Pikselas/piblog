function GetBlogTile(title: string , description: string, tags: string[])
{
    let blog_tile = document.createElement("div");
    blog_tile.className = "blog_itm";

    let blog_background_style = document.createElement("div");

    let hex_clover = new HexClover(450 , 550);
    
    AsElement(hex_clover).style.position = "absolute";
    AsElement(hex_clover).style.left = "315px";
    AsElement(hex_clover).style.top = "-166px";
    SetFlowerTheme(hex_clover, "sci_fi");

    blog_background_style.appendChild(AsElement(hex_clover).HtmlElement);

    let hex_clover2 = new HexClover(300 , 300);

    AsElement(hex_clover2).style.position = "absolute";
    AsElement(hex_clover2).style.left = "-115px";
    AsElement(hex_clover2).style.top = "-25px";
    AsElement(hex_clover2).style.transform = "rotate(90deg)";
    SetFlowerTheme(hex_clover2, "sci_fi");

    blog_background_style.appendChild(AsElement(hex_clover2).HtmlElement);

    blog_tile.appendChild(blog_background_style);

    let content_area = document.createElement("div");
    content_area.className = "content_area";

    let content = document.createElement("div");
    content.className = "content";

    let title_element = document.createElement("h1");
    title_element.innerText = title;

    content.appendChild(title_element);

    let description_element = document.createElement("div");
    description_element.className = "description";
    description_element.innerText = description;

    content.appendChild(description_element);

    let tags_element = document.createElement("div");
    tags_element.className = "tags";

    tags.forEach(tag => 
    {
        let tag_element = document.createElement("img");
        tag_element.className = "tag";
        tag_element.src = "./media/tag_" + tag + ".png";
        tags_element.appendChild(tag_element);
    });
    
    content.appendChild(tags_element);
    content_area.appendChild(content);

    let base_img = document.createElement("img");
    base_img.src = "./media/lowpoly_planet.png";
    base_img.className = "blog_pan_img";

    content_area.appendChild(base_img);
    blog_tile.appendChild(content_area);

    return blog_tile;
}

function CreateTextBlogItem(text: string)
{
    let blog_tile = document.createElement("p");
    blog_tile.innerHTML = text;
    return blog_tile;
}

function CreateCodeBlogItem(code: string)
{
    let blog_tile = document.createElement("pre");
    blog_tile.classList.add("code_snippet");
    blog_tile.innerHTML = code;
    return blog_tile;
}

function CreateImageBlogItem(src: string)
{
    let blog_tile = document.createElement("div");
    blog_tile.className = "img_container";

    let img = document.createElement("img");
    img.src = src;
    
    blog_tile.appendChild(img);
    return blog_tile;
}