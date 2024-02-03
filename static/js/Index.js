import Post from "./components/Post.js";

customElements.define("post-component", Post);

window.addEventListener("load", () => {
    const main = document.getElementById("main");
    const button = document.createElement("button");
    button.innerText = "Click me!";
    button.type = "button";
    button.addEventListener("click", () => {
        let posts = document.getElementsByTagName("post-component");
        main.removeChild(posts[0]);
    })
    main.appendChild(button);
    for (let i = 0; i < 10; i++) {
        let data = {
            title: `POST ${i+1}`,
            body: `BODY ${i+1}`
        };
        const p = createPost(data);
        main.appendChild(p);
    }
})

function createPost(data) {
    const p = document.createElement("post-component");
    p.innerHTML = `
        <span slot="post-title">${data.title}</span>
        <span slot="post-body">${data.body}</span>
    `;
    return p;
}