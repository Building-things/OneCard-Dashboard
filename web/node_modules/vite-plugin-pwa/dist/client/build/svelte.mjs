// src/client/build/svelte.ts
import { writable } from "svelte/store";

// src/client/build/register.ts
var autoUpdateMode = "__SW_AUTO_UPDATE__";
var selfDestroying = "__SW_SELF_DESTROYING__";
var auto = autoUpdateMode === "true";
var autoDestroy = selfDestroying === "true";
function registerSW(options = {}) {
  const {
    immediate = false,
    onNeedRefresh,
    onOfflineReady,
    onRegistered,
    onRegisteredSW,
    onRegisterError
  } = options;
  let wb;
  let registration;
  let registerPromise;
  let sendSkipWaitingMessage;
  const updateServiceWorker = async (reloadPage = true) => {
    await registerPromise;
    if (!auto) {
      if (reloadPage) {
        wb == null ? void 0 : wb.addEventListener("controlling", (event) => {
          if (event.isUpdate)
            window.location.reload();
        });
      }
      await (sendSkipWaitingMessage == null ? void 0 : sendSkipWaitingMessage());
    }
  };
  async function register() {
    if ("serviceWorker" in navigator) {
      const { Workbox, messageSW } = await import("workbox-window");
      sendSkipWaitingMessage = async () => {
        if (registration && registration.waiting) {
          await messageSW(registration.waiting, { type: "SKIP_WAITING" });
        }
      };
      wb = new Workbox("__SW__", { scope: "__SCOPE__", type: "__TYPE__" });
      wb.addEventListener("activated", (event) => {
        if (event.isUpdate)
          auto && window.location.reload();
        else if (!autoDestroy)
          onOfflineReady == null ? void 0 : onOfflineReady();
      });
      if (!auto) {
        const showSkipWaitingPrompt = () => {
          onNeedRefresh == null ? void 0 : onNeedRefresh();
        };
        wb.addEventListener("waiting", showSkipWaitingPrompt);
        wb.addEventListener("externalwaiting", showSkipWaitingPrompt);
      }
      wb.register({ immediate }).then((r) => {
        registration = r;
        if (onRegisteredSW)
          onRegisteredSW("__SW__", r);
        else
          onRegistered == null ? void 0 : onRegistered(r);
      }).catch((e) => {
        onRegisterError == null ? void 0 : onRegisterError(e);
      });
    }
  }
  registerPromise = register();
  return updateServiceWorker;
}

// src/client/build/svelte.ts
function useRegisterSW(options = {}) {
  const {
    immediate = true,
    onNeedRefresh,
    onOfflineReady,
    onRegistered,
    onRegisteredSW,
    onRegisterError
  } = options;
  const needRefresh = writable(false);
  const offlineReady = writable(false);
  const updateServiceWorker = registerSW({
    immediate,
    onOfflineReady() {
      offlineReady.set(true);
      onOfflineReady == null ? void 0 : onOfflineReady();
    },
    onNeedRefresh() {
      needRefresh.set(true);
      onNeedRefresh == null ? void 0 : onNeedRefresh();
    },
    onRegistered,
    onRegisteredSW,
    onRegisterError
  });
  return {
    needRefresh,
    offlineReady,
    updateServiceWorker
  };
}
export {
  useRegisterSW
};
