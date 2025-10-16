imgArray = [
    "static/img/hand_rock.png",
    "static/img/hand_paper.png",
    "static/img/hand_scissors.png"
];

// Función para reiniciar la UI cuando se inicia un nuevo juego
function resetGameUI() {
    // Ocultar la sección de resultados con animación
    const resultsSection = document.getElementById("results-section");
    if (resultsSection) {
        resultsSection.style.opacity = "0";
        setTimeout(() => {
            resultsSection.style.display = "none";
        }, 500);
    }
    
    // Resetear imágenes a la imagen inicial
    const imgPlayer = document.getElementById("imgPlayer");
    const imgComputer = document.getElementById("imgComputer");
    
    if (imgPlayer) {
        imgPlayer.src = "static/img/reload_icon.png";
    }
    if (imgComputer) {
        imgComputer.src = "static/img/reload_icon.png";
    }
}

// Ejecutar al cargar la página para asegurar el estado inicial
document.addEventListener('DOMContentLoaded', function() {
    resetGameUI();
});

function choose(option) {
    fetch(`/play?choice=${option}`)
        .then(response => response.json())
        .then(data => {
            // Mostrar la sección de resultados después de la primera jugada con animación
            const resultsSection = document.getElementById("results-section");
            if (resultsSection.style.display === "none") {
                resultsSection.style.display = "block";
                setTimeout(() => {
                    resultsSection.style.opacity = "1";
                }, 10);
            }
            
            // Update the UI with the results
            if (option === 0) {
                document.getElementById("player_choice").innerHTML = "El jugador eligió roca";
                document.getElementById("imgPlayer").src = "static/img/hand_rock.png";
            } else if (option === 1) {
                document.getElementById("player_choice").innerHTML = "El jugador eligió papel";
                document.getElementById("imgPlayer").src = "static/img/hand_paper.png";
            } else if (option === 2) {
                document.getElementById("player_choice").innerHTML = "El jugador eligió tijera";
                document.getElementById("imgPlayer").src = "static/img/hand_scissors.png";
            }

            // Actualizar scores
            document.getElementById("player_score").innerText = `Puntuación: ${data.player_score}`;
            document.getElementById("computer_score").innerText = `Puntuación: ${data.computer_score}`;
            
            // Actualizar elección de la computadora
            document.getElementById("computer_choice").innerText = data.computer_choice;
            
            // Actualizar resultado con estilo dinámico
            const roundResultElement = document.getElementById("round_result");
            roundResultElement.innerText = data.round_result;
            
            // Aplicar estilos según el resultado
            roundResultElement.className = "list-group-item"; // Reset classes
            if (data.round_result === "Ganaste") {
                roundResultElement.classList.add("text-success", "fw-bold");
            } else if (data.round_result === "Perdiste") {
                roundResultElement.classList.add("text-danger", "fw-bold");
            } else {
                roundResultElement.classList.add("text-warning", "fw-bold");
            }
            
            // Actualizar mensaje
            const messageElement = document.getElementById("message");
            messageElement.innerText = data.message;
            messageElement.className = "list-group-item"; // Reset classes
            if (data.round_result === "Ganaste") {
                messageElement.classList.add("text-success");
            } else if (data.round_result === "Perdiste") {
                messageElement.classList.add("text-danger");
            } else {
                messageElement.classList.add("text-info");
            }

            // Actualizar imagen de la computadora
            var imgComputer = document.getElementById("imgComputer");
            imgComputer.src = imgArray[data.computer_choice_index];
        })
        .catch(error => console.error("Error:", error));
}
