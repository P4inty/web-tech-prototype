<script>
    import { Views } from "../enums/views";
    import { ViewManager } from "../stores/view";

    const { view } = ViewManager;
</script>

<div class="flex w-full space-x-4 items-stretch px-4">
    <div class="w-1/2 md:w-1/4 pt-4">
        <button
            class="nav-btn"
            class:active={$view === Views.UPLOAD_FILE}
            on:click={() => ViewManager.goTo(Views.UPLOAD_FILE)}
        >
            <span>Upload File</span>
            <span class="text-xs block mt-4">Upload all your files here.</span>
        </button>
    </div>
    <div class="w-1/2 md:w-1/4 pt-4">
        <button
            class="nav-btn"
            class:active={$view === Views.FILES}
            on:click={() => ViewManager.goTo(Views.FILES)}
        >
            <span>Files</span>
            <span class="text-xs block mt-4"
                >See all of your favorite uploaded files here.</span
            >
        </button>
    </div>
</div>
<hr style="margin-top: -1px;" class="border-primary" />
