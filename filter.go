package main

import (
	"fmt"
	"strings"
)

func FilteredVerboseOutput(verbose bool, text string) {
	if verbose == true {
		fmt.Println(text)
	} else {
		return
	}
}

func FilteredDataOutput(s []string) []string {
	var blacklists = []string{
		"node_modules", "jquery", "bootstrap", "react", "vue", "angular", "favicon.ico", "logo", "style.css",
		"font-awesome", "materialize", "semantic-ui", "tailwindcss", "bulma", "d3", "chart.js", "three.js",
		"vuex", "express", "axios", "jquery.min.js", "moment.js", "underscore", "lodash", "jquery-ui",
		"angular.min.js", "react-dom", "redux", "chartist.js", "anime.min.js", "highcharts", "leaflet",
		"pdf.js", "fullcalendar", "webfontloader", "swiper", "slick.js", "datatables", "webfonts", "react-scripts",
		"vue-router", "vite", "webpack", "electron", "socket.io", "codemirror", "angularjs", "firebase", "swagger",
		"typescript", "p5.js", "ckeditor", "codemirror.js", "recharts", "bluebird", "lodash.min.js", "sweetalert2",
		"polyfils", "runtime", "bootstrap", "google-analytics",
		"application/json", "application/x-www-form-urlencoded", "json2.js", "querystring", "axios.min.js",
		"ajax", "formdata", "jsonschema", "jsonlint", "json5", "csrf", "jQuery.ajax", "superagent",
		"body-parser", "urlencoded", "csrf-token", "express-session", "content-type", "fetch", "protobuf",
		"formidable", "postman", "swagger-ui", "rest-client", "swagger-axios", "graphql", "apollo-client",
		"react-query", "jsonapi", "json-patch", "urlencoded-form", "url-search-params", "graphql-tag",
		"vue-resource", "graphql-request", "restful-api", "jsonwebtoken", "fetch-jsonp", "reqwest", "lodash-es",
		"jsonwebtoken", "graphene", "axios-jsonp", "postman-collection",
		"application/xml", "text/xml", "text/html", "text/plain", "multipart/form-data", "image/jpeg",
		"image/png", "image/gif", "audio/mpeg", "audio/ogg", "video/mp4", "video/webm", "text/css",
		"application/pdf", "application/octet-stream", "image/svg+xml", "application/javascript",
		"application/ld+json", "text/javascript", "application/x-www-form-urlencoded", ".dtd", "google.com", "application/javascript", "text/css", "w3.org", "www.thymeleaf.org", "application/javascrip", "toastr.min.js", "spin.min.js", "./", "DD/MM/YYYY",
	}

	filtered := make(map[string]bool)

	var results []string

	for _, dataOut := range s {
		if filtered[dataOut] {
			continue
		}
		filtered[dataOut] = true

		skip := false

		for _, blacklist := range blacklists {
			if strings.Contains(dataOut, blacklist) {
				skip = true
				break
			}
		}

		if skip {
			continue
		}

		dataOut = strings.ReplaceAll(dataOut, `"`, "")
		dataOut = strings.ReplaceAll(dataOut, `'`, "")

		results = append(results, dataOut)
	}

	return results
}
