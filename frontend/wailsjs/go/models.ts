export namespace reference {
	
	export class Verse {
	    Number: number;
	    Text: string;
	
	    static createFrom(source: any = {}) {
	        return new Verse(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.Number = source["Number"];
	        this.Text = source["Text"];
	    }
	}
	export class BiblePassage {
	    Book: string;
	    BookId?: number;
	    Chapter: number;
	    StartVerse: number;
	    EndVerse: number;
	    FullText: Verse[];
	
	    static createFrom(source: any = {}) {
	        return new BiblePassage(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.Book = source["Book"];
	        this.BookId = source["BookId"];
	        this.Chapter = source["Chapter"];
	        this.StartVerse = source["StartVerse"];
	        this.EndVerse = source["EndVerse"];
	        this.FullText = this.convertValues(source["FullText"], Verse);
	    }
	
		convertValues(a: any, classs: any, asMap: boolean = false): any {
		    if (!a) {
		        return a;
		    }
		    if (a.slice && a.map) {
		        return (a as any[]).map(elem => this.convertValues(elem, classs));
		    } else if ("object" === typeof a) {
		        if (asMap) {
		            for (const key of Object.keys(a)) {
		                a[key] = new classs(a[key]);
		            }
		            return a;
		        }
		        return new classs(a);
		    }
		    return a;
		}
	}
	export class BibleReference {
	    Passages: BiblePassage[];
	
	    static createFrom(source: any = {}) {
	        return new BibleReference(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.Passages = this.convertValues(source["Passages"], BiblePassage);
	    }
	
		convertValues(a: any, classs: any, asMap: boolean = false): any {
		    if (!a) {
		        return a;
		    }
		    if (a.slice && a.map) {
		        return (a as any[]).map(elem => this.convertValues(elem, classs));
		    } else if ("object" === typeof a) {
		        if (asMap) {
		            for (const key of Object.keys(a)) {
		                a[key] = new classs(a[key]);
		            }
		            return a;
		        }
		        return new classs(a);
		    }
		    return a;
		}
	}

}

