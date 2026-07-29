function partOne(): void {
    require("fs").readFile("input.txt", "utf-8", (err: Error, data: string) => {
        if (err) {
            console.error("Error reading file:", err);
            return;
        }

        const lines: string[] = data.split("\n");
        const emptyIndex: number = lines.findIndex(line => line.trim() === "");
        const linesFirstPart: string[] = lines.slice(0, emptyIndex).filter(line => line.trim() !== "");
        const linesSecondPart: string[] = lines.slice(emptyIndex + 1).filter(line => line.trim() !== "");
        var count = 0;

        for(let i = 0; i < linesSecondPart.length; i++) {
            const number: number = parseInt(linesSecondPart[i]);
            for(let j = 0; j < linesFirstPart.length; j++) {
                const from: number = parseInt(linesFirstPart[j].split("-")[0]);
                const to: number = parseInt(linesFirstPart[j].split("-")[1]);
                if(number >= from && number <= to) {
                    count++;
                    break;
                }
            }   
        }
        console.log("Count:", count);
    });
}

function partTwo(): void {
    require("fs").readFile("input.txt", "utf-8", (err: Error, data: string) => {
        if (err) {
            console.error("Error reading file:", err);
            return;
        }

        const lines: string[] = data.split("\n");
        const emptyIndex: number = lines.findIndex(line => line.trim() === "");
        const linesFirstPart: string[] = lines.slice(0, emptyIndex).filter(line => line.trim() !== "");
        const linesSecondPart: string[] = lines.slice(emptyIndex + 1).filter(line => line.trim() !== "");
        var sum = 0;
        
        const intervals = linesFirstPart.map(line => {
            const parts = line.split("-");
            return { 
                from: parseInt(parts[0], 10), 
                to: parseInt(parts[1], 10) 
            };
        }).sort((a, b) => a.from - b.from);

        let currentFrom = intervals[0].from;
        let currentTo = intervals[0].to;

        // const numbersUsed = new Set<number>();

        for (let i = 1; i < intervals.length; i++) {
            const next = intervals[i];
            if (next.from <= currentTo) {
                if (next.to > currentTo) {
                    currentTo = next.to;
                }
            } else {
                sum += (currentTo - currentFrom + 1);
                currentFrom = next.from;
                currentTo = next.to;
            }
        }

        sum += (currentTo - currentFrom + 1);

        // for(let i = 0; i < linesFirstPart.length; i++) {
        //     const from: number = parseInt(linesFirstPart[i].split("-")[0]);
        //     const to: number = parseInt(linesFirstPart[i].split("-")[1]);
        //     for(let j = from; j <= to; j++) {
        //         if(!numbersUsed.has(j)) {
        //             numbersUsed.add(j);
        //             sum++;
        //         }
        //     }
        // }

        console.log("Sum fresh:", sum);
    });
}
partOne();
partTwo();
