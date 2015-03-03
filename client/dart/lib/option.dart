part of socket;

class Option {
    
    String method;
    String url;
    String data;
    int timeout;
    
    Option({String method: "GET", String url : "", String data : "{}", int timeout : 60}) {
        this.method = method;
        this.url = url;
        this.data = data;
        this.timeout = timeout;
    }    
    
    String get Method  => method;
    String get Url     => url;
    String get Data    => data;
    int    get Timeout => timeout;
    
    void set Method(String method) {
        this.Method = method;
    }
    
    void set Url(String url) {
        this.url = url;
    }
    
    void set Data(String data) {
        this.data = data;
    }
    
    void set Timeout(int timeout) {
        this.timeout = timeout;
    }
    
}