document.getElementById("loginForm").addEventListener("submit", async function(event) {
    event.preventDefault();

    const usuario = document.getElementById("usuario").value;
    const contrasena = document.getElementById("contrasena").value;

    const response = await fetch("http://localhost:8080/api/login", {
        method: "POST",
        headers: { "Content-Type": "application/json" },
        body: JSON.stringify({ usuario, contrasena })
    });

    if (response.ok) {
        window.location.href = "inicio.html";  // Redirige si el login es exitoso
    } else {
        document.getElementById("errorMsg").style.display = "block";
        setTimeout(() => { window.location.href = "error404.html"; }, 2000);
    }
});
