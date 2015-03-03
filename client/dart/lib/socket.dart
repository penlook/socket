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

library socket;

import "dart:html";

part "transport.dart";
part "polling.dart";
part "event.dart";
part "option.dart";

/**
 * Socket
 *
 * @category   Socket
 * @package    Service
 * @copyright  Penlook Development Team
 * @license    GNU Affero General Public
 * @version    1.0
 * @link       http://github.com/penlook
 * @since      Class available since Release 1.0
 */
class Socket {

    /**
     * Socket protocol
     *
     * @var string http | https
     */
    String protocol;

    /**
     * Host name
     *
     * @var string
     */
    String host;

    /**
     * Server port
     *
     * @var int
     */
    int port;

    /**
     * Socket server url
     *
     * @var string
     */
    String url;

    /**
     * Socket contructor
     *
     * @param string protocol
     * @param string localhost
     * @param int    port 80
     */
    Socket({
        String protocol : "http",
        String host : "localhost",
        int port: 80
    }) {
        this.host = host;
        this.port = port;
        this.protocol = "http";
        this.url = protocol + "://" + host + ":" + port.toString();
    }

    String get Protocol => this.protocol;
    String get Url      => this.url;
    String get Host     => this.host;
    int    get Port     => this.port;

    syncRequest(Socket context, Option option, Function callback) {

        // Initialize new HTTP Request
        HttpRequest request = new HttpRequest();
        request.open(option.Method, option.Url, async: option.Async);
        request.setRequestHeader("Content-Type", "application/json;charset=UTF-8");
        request.send(option.Data);
    }

    asyncRequest(Socket context, Option option, Function callback) {

    }

    processResponse() {

    }

    on() {

    }

    emit() {

    }

    remove() {

    }

    connect() {

    }

    pull() {

    }

    push() {

    }

}
