/** @module writer-panel-util */

import $ from "jquery";
import "select2";

/**
 * WriterPanelUtil contains utility functions for writer-panel pages
 */
class WriterPanelUtil {

	// Initializes the tag selector
	initTagSelector() {
		$(".tag-selector").select2({
			ajax: {
				url:      "/api/v1/search",
				dataType: "json",
				delay:    250,
				cache:    true,
				width:    "100%",
				data:     params => {
					return {
						q: "type:tag " + params.term,
					};
				},
				processResults: data => {
					if(data) {
						return {
							results: data.Tags.map(e => {
								return {
									id:          e.ID,
									name:        e.Name,
									description: e.Description,
								};
							}),
						};
					}
					return {results: []};
				},
			},
			placeholder:    "Type to search for tags...",
			allowClear:     true,
			templateResult: data => {
				return data.name + " - " + data.description;
			},
			templateSelection: data => {
				if(!data.text)
					return data.name;
				return data.text;
			},
		});
	}
}

export default WriterPanelUtil;
