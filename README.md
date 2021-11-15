# ApartmentLayoutReader
A golang project to read the blueprint of an apartment and identify the placements of different chair in different tooms

The layout of an apartment may look as follows :
<pre>+-----------+------------------------------------+
|           |                                    |
| (closet)  |                                    |
|         P |                            S       |
|         P |         (sleeping room)            |
|         P |                                    |
|           |                                    |
+-----------+    W                               |
|           |                                    |
|        W  |                                    |
|           |                                    |
|           +--------------+---------------------+
|                          |                     |
|                          |                W W  |
|                          |    (office)         |
|                          |                     |
+--------------+           |                     |
|              |           |                     |
| (toilet)     |           |             P       |
|   C          |           |                     |
|              |           |                     |
+--------------+           +---------------------+
|              |           |                     |
|              |           |                     |
|              |           |                     |
| (bathroom)   |           |      (kitchen)      |
|              |           |                     |
|              |           |      W   W          |
|              |           |      W   W          |
|       P      +           |                     |
|             /            +---------------------+
|            /                                   |
|           /                                    |
|          /                          W    W   W |
+---------+                                      |
|                                                |
| S                                   W    W   W |
|                (living room)                   |
| S                                              |
|                                                |
|                                                |
|                                                |
|                                                |
+--------------------------+---------------------+
                           |                     |
                           |                  P  |
                           |  (balcony)          |
                           |                 P   |
                           |                     |
                           +---------------------+</pre>

# How to get started

```
./start.sh

go build
go test -v
```

# Things that needs to be done
We now need a command line tool that reads in such a file and outputs the following information:
- Number of different chair types for the apartment
- Number of different chair types per room

The different types of chairs are as follows:<br/>
W: wooden chair<br/>
P: plastic chair<br/>
S: sofa chair<br/>
C: china chair<br/>

The output must look like as follows:<br/>

total:<br/>
W: 3, P: 2, S: 0, C: 0<br/>
living room:<br/>
W: 3, P: 0, S: 0, C: 0<br/>
office:<br/>
W: 0, P: 2, S: 0, C: 0<br/>