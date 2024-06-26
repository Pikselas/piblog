
function ShowBlogPosts(blog_data: any)
{
    let post_container = document.getElementById("blog_posts");
    post_container.innerHTML = "";
    blog_data.forEach((blog) => {
        let ref = document.createElement("a");
        ref.href = "/blog/" + blog.Id;
        ref.appendChild(GetBlogTile(blog.Title, blog.Description, blog.Tags));
        post_container.appendChild(ref);
    })
    document.getElementById("blog_itm_container").style.top = "calc(50% - 42.5%)";
}

async function ShowBlogPostsByTag(query: string)
{
    let blogs = await fetch("/search_by_tags",{method: "POST",body: JSON.stringify([query])});
    let blog_data = await blogs.json();
    ShowBlogPosts(blog_data);
}

async function ShowBlogPostsBySearch(query: string)
{
    let blogs = await fetch("/search_by_title/" + query);
    let blog_data = await blogs.json();
    ShowBlogPosts(blog_data);
}

function HideBlogPosts()
{
    document.getElementById("blog_itm_container").style.top = "100%";
}