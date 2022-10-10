<script lang="ts">
    export let file: File;

    function handleFile(e: any) {
        const files: FileList = e.target.files;
        if (files.length === 1) {
            file = files.item(0);
        }
    }
</script>

<div class="flex justify-center items-center w-full mt-4">
    <label
        for="dropzone-file"
        class="flex flex-col justify-center items-center w-full h-full rounded-lg border-2 border-primary border-dashed cursor-pointer hover:bg-primary hover:bg-opacity-20"
    >
        <div class="flex flex-col justify-center items-center pt-5 pb-6">
            <svg
                aria-hidden="true"
                class="mb-3 w-10 h-10"
                fill="none"
                stroke="currentColor"
                viewBox="0 0 24 24"
                xmlns="http://www.w3.org/2000/svg"
                ><path
                    stroke-linecap="round"
                    stroke-linejoin="round"
                    stroke-width="2"
                    d="M7 16a4 4 0 01-.88-7.903A5 5 0 1115.9 6L16 6a5 5 0 011 9.9M15 13l-3-3m0 0l-3 3m3-3v12"
                /></svg
            >
            <p class="mb-2 text-sm">Click to upload or drag and drop</p>
            <p class="text-xs">All file types supported</p>
        </div>
        <input
            id="dropzone-file"
            type="file"
            class="hidden"
            on:change={(e) => handleFile(e)}
        />
    </label>
</div>
