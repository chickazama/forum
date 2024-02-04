const content = `
<style>
    .test {
        color: red;
    }
    .test:hover {
        color: blue;
    }
</style>
<div class="post">
    <h1><slot name="post-category" class="test">CATEGORY</slot></h1>
    <h3><slot name="post-content">CONTENT</slot></h3>
    <h5>Likes: <slot name="post-likes">LIKES</slot></h5>
    <h5>Dislikes: <slot name="post-dislikes">DISLIKES</slot></h5>
</div>
`;

export default class Post extends HTMLElement {
    shadowRoot;
    template;
    constructor() {
        super();
        this.shadowRoot = this.attachShadow({mode: "open"});
        this.template = document.createElement("template");
        this.template.innerHTML = content;
        this.shadowRoot.appendChild(this.template.content.cloneNode(true));
    }

    connectedCallback() {
        console.log("Post component added to DOM.");
    }

    disconnectedCallback() {
        console.log("Post component removed from DOM.");
    }
}