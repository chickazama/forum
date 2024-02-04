const content = `
<style>
    .test {
        color: red;
    }
    .test:hover {
        color: blue;
    }
    .post {
        background-color: green;
    }
</style>
<div class="post">
    <h1><slot name="post-title" class="test">TITLE</slot></h1>
    <h3><slot name="post-body">BODY></slot></h3>
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