import Post from "./components/Post.js";

customElements.define("post-component", Post);

window.addEventListener("load", async () => {
    const main = document.getElementById("main");
    const btn = document.createElement("button");
    btn.type = "button";
    btn.innerText = "Click Me!";
    btn.addEventListener("click", ()=> {
        removePostByID(1);
    })
    main.appendChild(btn);
    const res = await fetch("http://localhost:7777/posts");
    const body = await res.json();
    for (const data of body) {
        console.log(data);
        const post = createPost(data);
        main.appendChild(post);
    }
})

function createPost(data) {
    const p = document.createElement("post-component");
    p.id = `post-${data.PostID}`;
    p.innerHTML = `
        <span slot="post-category">${data.Category}</span>
        <span slot="post-content">${data.Content}</span>
    `;
    return p;
}

function removePost(data) {
    const main = document.getElementById("main");
    const p = document.getElementById(`post-${data.PostID}`);
    if (!p) {
        return;
    }
    main.removeChild(p);
}

function removePostByID(id) {
    const main = document.getElementById("main");
    const p = document.getElementById(`post-${id}`);
    if (!p) {
        return;
    }
    main.removeChild(p);
}