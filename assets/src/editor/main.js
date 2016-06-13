import Editor from "./editor.js";
import {DebugSaveManager} from "./save-manager.js";
import {SlowSaveManager} from "./debug.js";

const saveManager = new SlowSaveManager(new DebugSaveManager());
const editor = new Editor(document.getElementById("editor"), saveManager);

window.editor = editor;
export default editor;
