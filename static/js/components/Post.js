const content = `
<link href="./static/css/Post.css" rel="stylesheet" type="text/css" />
<h1><slot name="post-title" class="test">TITLE</slot></h1>
<h3><slot name="post-body">BODY</h3>
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