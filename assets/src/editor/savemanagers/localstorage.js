import SaveManager from "../save-manager.js";

export default class LocalStorageSaveManager extends SaveManager {
	saveSync(document) {
		window.localStorage.setItem("project-domino-editor", JSON.stringify(document));
	}
}
