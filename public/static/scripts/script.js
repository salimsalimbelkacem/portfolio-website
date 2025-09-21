function lamp(params) {
	lamp = document.createElement("img")
	lamp.src="/static/images/lamp.png"

	lamp_hover = "top-[20%] left-[80%] w-7 fixed cursor-grab"
	lamp_grab =  "top-[20%] left-[80%] w-7 fixed cursor-grabbing"

	lamp.className=lamp_hover

	follower = document.createElement("div")
	follower.className="top-[20%] left-[80%] bg-yellow-100 fixed w-30 h-30 blur-2xl -z-10 rounded-full animate-lamp-light"

	follower.style.transform = "translate(-50%, -50%)";
	lamp.style.transform = "translate(-50%, -50%)";

	document.body.appendChild(follower)
	document.body.appendChild(lamp)

	dragElement(lamp)

	function dragElement(elmnt) {
		let pos1 = 0, pos2 = 0, pos3 = 0, pos4 = 0;
		elmnt.onmousedown = dragMouseDown;

		function dragMouseDown(e) {
			e = e || window.event;
			e.preventDefault();
			pos3 = e.clientX;
			pos4 = e.clientY;
			elmnt.className=lamp_grab
			document.onmouseup = closeDragElement;
			document.onmousemove = elementDrag;
		}

		function elementDrag(e) {
			e = e || window.event;
			e.preventDefault();
			pos1 = pos3 - e.clientX;
			pos2 = pos4 - e.clientY;
			pos3 = e.clientX;
			pos4 = e.clientY;
			elmnt.style.top = (elmnt.offsetTop - pos2) + "px";
			elmnt.style.left = (elmnt.offsetLeft - pos1) + "px";

			follower.style.top =  (elmnt.offsetTop - pos2) + "px";
			follower.style.left = (elmnt.offsetLeft - pos1) + "px";

		}

		function closeDragElement() {
			document.onmouseup = null;
			document.onmousemove = null;
			elmnt.className=lamp_hover
		}
	}

}

window.isDesktop = window.matchMedia("(min-width: 768px)").matches;

if(window.isDesktop) {
    lamp();
}
