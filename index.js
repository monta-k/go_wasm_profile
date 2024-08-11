import "./wasm_exec.js";
import module from "./build/app.wasm";

export default {
  async fetch(request, env, ctx) {
    const go = new Go();
    const result = await WebAssembly.instantiate(module, go.importObject);
    go.run(result);
    const response = serveHTML();
    return new Response(response, {
      headers: {
        "content-type": "text/html;charset=UTF-8",
      },
    });
  },
};
