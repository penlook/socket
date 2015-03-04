/**
 * Penlook Project
 *
 * Copyright (c) 2015 Penlook Development Team
 *
 * --------------------------------------------------------------------
 *
 * This program is free software: you can redistribute it and/or
 * modify it under the terms of the GNU Affero General Public License
 * as published by the Free Software Foundation, either version 3
 * of the License, or (at your option) any later version.
 *
 * This program is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
 * GNU Affero General Public License for more details.
 *
 * You should have received a copy of the GNU Affero General Public
 * License along with this program.
 * If not, see <http://www.gnu.org/licenses/>.
 *
 * --------------------------------------------------------------------
 *
 * Authors:
 *     Loi Nguyen       <loint@penlook.com>
 *     Nam Vo           <namvh@penlook.com>
 */

part of socket.polling;

abstract class Transport {    

    /**
     * HTTP Synchronous request
     *
     * @param Socket context
     * @param Option option
     * @param Function callback
     */
    void syncRequest(Object context, Option option, Function callback(Object context, Map<String, Map> response)) {

        // Initialize new HTTP Request
        HttpRequest request = new HttpRequest();

        request.open(option.Method, option.Url, async: false);        

        Map<String, Map> response = null;

        try {
           request.send(option.Data);
           response = JSON.decode(request.responseText);
        } catch (e) {
           throw e;
       }

       callback(context, response);
    }

    /**
     * HTTP Asynchronous request
     *
     * @param Socket context
     * @param Option option
     * @param Function callback
     */
    void asyncRequest(Object context, Option option, Function callback(Object context, Map<String, Map> response)) {
    
    }

    void sendRequest(Object context, Option option, Function callback) {
        option.Async ?
            this.asyncRequest(context, option, callback) :
                this.syncRequest(context, option, callback);
    }
    
}