body = document.body;
follower = document.createElement("div")

follower.className="hidden md:inline bg-white fixed w-15 h-15 blur-xl -z-10 rounded-full"

body.appendChild(follower)

document.addEventListener("mousemove", (e) => {
  follower.style.transform = `translate(${e.clientX - 30}px, ${e.clientY - 30}px)`;
});
