export namespace main {
	
	export class Clock {
	    timezone: string;
	    name: string;
	    timeFormat: string;
	    dateFormat: string;
	    color: string;
	
	    static createFrom(source: any = {}) {
	        return new Clock(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.timezone = source["timezone"];
	        this.name = source["name"];
	        this.timeFormat = source["timeFormat"];
	        this.dateFormat = source["dateFormat"];
	        this.color = source["color"];
	    }
	}

}

