(() => {
  const passwordEl = document.getElementById("password");
  const genBtn = document.getElementById("genBtn");
  const copyBtn = document.getElementById("copyBtn");

  genBtn.addEventListener("click", async () => {
    genBtn.disabled = true;

    try {
      const res = await fetch("/api/generate");
      if (!res.ok) throw new Error("Falha ao gerar senha");

      const data = await res.json();
      passwordEl.value = data.password || "";
      copyBtn.disabled = !passwordEl.value;

      passwordEl.style.boxShadow = "0 0 15px #00ff88";
      setTimeout(() => {
        passwordEl.style.boxShadow = "none";
      }, 300);
    } catch (e) {
      copyBtn.disabled = true;
    } finally {
      genBtn.disabled = false;
    }
  });

  copyBtn.addEventListener("click", async () => {
    try {
      await navigator.clipboard.writeText(passwordEl.value);

      copyBtn.style.background = "#00ff88";
      copyBtn.style.color = "#000";
      setTimeout(() => {
        copyBtn.style.background = "";
        copyBtn.style.color = "";
      }, 200);
    } catch (e) {}
  });
})();
