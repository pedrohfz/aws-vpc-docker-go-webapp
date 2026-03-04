(() => {
  const passwordEl = document.getElementById("password");
  const genBtn = document.getElementById("genBtn");
  const copyBtn = document.getElementById("copyBtn");
  const msgEl = document.getElementById("msg");

  const setMsg = (text) => {
    msgEl.textContent = text || "";
  };

  genBtn.addEventListener("click", async () => {
    setMsg("");
    genBtn.disabled = true;

    try {
      const res = await fetch(`/api/generate`);
      if (!res.ok) throw new Error("Falha ao gerar senha");

      const data = await res.json();
      passwordEl.value = data.password || "";
      copyBtn.disabled = !passwordEl.value;

      setMsg("Senha gerada com sucesso ✅");
    } catch (e) {
      setMsg(
        "Não foi possível gerar a senha. Verifique se o backend está rodando.",
      );
      copyBtn.disabled = true;
    } finally {
      genBtn.disabled = false;
    }
  });

  copyBtn.addEventListener("click", async () => {
    try {
      await navigator.clipboard.writeText(passwordEl.value);
      setMsg("Copiado! 📋");
    } catch (e) {
      setMsg("Não consegui copiar automaticamente. Copie manualmente.");
    }
  });
})();
