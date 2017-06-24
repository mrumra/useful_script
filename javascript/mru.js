//获取http请求参数键值对
function getQueryStringArgs() {
	var qs = (document.location.search.length > 0 ? document.location.search.substring(1) : "");
	var args = {};
	var items = qs.length ? qs.split("&") : [] 
	var item = null;
	var name = null;
	var value = null;
	var i = 0;
	var len = items.length;

	for (i = 0; i < len; i++) {
		item = items[i].split("=");
		name = decodeURIComponent(item[0]);
		value = decodeURIComponent(item[1]);

		if (name.length) {
			args[name] = value;
		}
	}
	return args;
}

//检查某个对象是否是可排序的
function isSortable(object) {
	return typeof object.sort == "function";
}

//检查某个对象是否有某个方法
function isHostMethod(object, property) {
	var t = typeof object[property];
	return t == "function" ||
		(!!(t=="object" && object[property])) ||
		t == "unknown";
}

function getInnerText(element) {
	return (typeof element.textContent == "string") ? 
		element.textContent : element.innertText;
}

function setInnerText(element, text) {
	if (typeof element.textContext == "string") {
		element.textContent == text;
	} else {
		element.innertText = text;
	}
}

var CookieUtil = {
	get: function(name) {
		var cookieName = encodeURIComponent(name) + "=", 
			cookieStart = document.cookie.indexOf(cookieName),
			cookieValue = null;

		if (cookieStart > -1) {
			var cookieEnd = document.cookie.indexOf(";", cookieStart);
			if (cookieEnd == -1) {
				cookieEnd = document.cookie.length;
			}
			cookieValue = decodeURIComponent(document.cookie.substring(cookieStart + cookieName.length, cookieEnd));
		}

		return cookieValue;
	},

	set: function(name, value, expires, path, domain, secure) {
		var cookieText = encodeURIComponent(name) + "=" + encodeURIComponent(value);

		if (expires instanceof Date) {
			cookieText += "; expires=" + expires.toGMTString();
		}

		if (path) {
			cookieText += "; path=" + path;
		}

		if (domain) {
			cookieText += "; domain=" + domain;
		}

		if (secure) {
			cookieText += "; secure";
		}

		document.cookie = cookieText;
	},

	unset:  function(name, path, domain, secure) {
		this.set(name, "", new Date(0), path, domain, secure);
	}
};

$().ready(function() {
	alert("My script loaded");
});

