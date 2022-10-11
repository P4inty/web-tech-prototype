<script lang="ts">
  import { Views } from "./enums/views";
  import Footer from "./lib/footer.svelte";
  import Nav from "./lib/nav.svelte";
  import Files from "./pages/files.svelte";
  import FileUpload from "./pages/fileUpload.svelte";
  import { ViewManager } from "./stores/view";

  const { view } = ViewManager;
</script>

<main class="bg-black text-primary h-screen max-h-screen overflow-y-auto">
  <Nav />
  {#if $view === Views.UPLOAD_FILE}
    <FileUpload />
  {:else if $view === Views.FILES}
    <Files />
  {/if}
  <Footer />
</main>
