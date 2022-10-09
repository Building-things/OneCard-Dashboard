<script>
    export let size;
    export let value_of_total;
    export let value_of_current;


    const strokeWidth = 4
    const radius = (size / 2 ) - (strokeWidth * 2);
    const circumference = 2 * Math.PI * radius;

    const percentFilled =  value_of_current / value_of_total
    let offset = circumference - (percentFilled) * circumference;
</script>

<svg height={size} width={size}>
    <circle
        stroke-width="{strokeWidth}"
        fill="transparent"
        stroke=var(--accent-colour)
        r={radius}
        cx={size / 2}
        cy={size / 2}
        stroke-dasharray={`${circumference} ${circumference}`}
        stroke-dashoffset={offset}
        />
</svg>

<style>
</style>