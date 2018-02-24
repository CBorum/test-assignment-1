package com.cborum;

/**
 * Class with bugs.
 *
 *
 */
public class MyArrayListWithBugs {

    int nextFree;
    private Object[] list;

    // Creates a new empty list
    public MyArrayListWithBugs() {
        list = new Object[5];
        nextFree = 0;
    }

    // Inserts object at the end of list
    public void add(Object o) {
        // check capacity
        if (list.length <= nextFree)
            list = getLongerList();

        list[nextFree] = o;
        nextFree++;
    }

    // Returns the number of objects in the list
    public int size() {
        return nextFree;
    }

    // Returns a reference to the object at position index
    // Throws IndexOutOfBoundsException
    public Object get(int index) {
        if(index <= 0 || nextFree < index)
            throw new IndexOutOfBoundsException("Error (get): Invalid index" +
                    index);

        return list[index];
    }

    // Inserts object at position index
    // Throws IndexOutOfBoundsException
    public void add(int index, Object o) {
        if(index < 0 || nextFree < index)
            throw new IndexOutOfBoundsException("Error (add): Invalid index" +
                    index);

        // check capacity
        if (list.length <= nextFree)
            list = getLongerList();

        // Shift elements upwards to make position index free
        // Start with last element and move backwards
        for (int i = nextFree-1; i > index; i--) {
            list[i] = list[i-1];
        }

        list[index] = o;
    }

    // Removes object at position index
    // Returns a reference to the removed object
    // Throws IndexOutOfBoundsException
    public Object remove(int index) {
        if (index < 0 || nextFree <= index)
            throw new IndexOutOfBoundsException("Error (remove): Invalid index"
                    + index);


        // Shift elements down to fill indexed position
        // Start with first element
        for (int i = index; i < nextFree-1; i++) {
            list[i] = list[i+1];
        }
        nextFree--;

        return list[index];

    }

    //============== private helper methods ==========
    // create a list with double capacity and
    // copy all elements to this
    private Object[] getLongerList() {
        Object[] tempList = new Object[list.length*2];
        for (int i=0; i< list.length;i++) {
            tempList[i] = list[i];
        }
        return tempList;
    }


}